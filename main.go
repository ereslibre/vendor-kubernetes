/**
 * Copyright 2019 Rafael Fernández López <ereslibre@ereslibre.es>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/filemode"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

type options struct {
	kubernetesPath string
	kubernetesTag  string
}

var cmdOptions options

func kubernetesSubprojects(tree *object.Tree) ([]string, error) {
	res := []string{}
	stagingTree, err := tree.Tree(path.Join("staging", "src", "k8s.io"))
	if err != nil {
		return res, err
	}
	stagingTreeWalker := object.NewTreeWalker(stagingTree, false, map[plumbing.Hash]bool{})
	defer stagingTreeWalker.Close()
	for {
		name, entry, err := stagingTreeWalker.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return res, err
		}
		if entry.Mode&filemode.Dir != 0 {
			res = append(res, name)
		}
	}
	return res, nil
}

func checkoutTree(repo *git.Repository, tag string) (*object.Commit, *object.Tree, error) {
	referenceName, err := repo.Reference(plumbing.NewTagReferenceName(tag), true)
	if err != nil {
		return nil, nil, err
	}
	tagObject, err := repo.TagObject(referenceName.Hash())
	if err != nil {
		return nil, nil, err
	}
	commitObject, err := tagObject.Commit()
	if err != nil {
		return nil, nil, err
	}
	tree, err := tagObject.Tree()
	return commitObject, tree, err
}

func retrieveOrCloneKubernetesSubproject(subproject, tag string) (*object.Commit, *object.Tree, error) {
	repo, err := retrieveOrCloneKubernetesSubprojectRepo(subproject)
	if err != nil {
		return nil, nil, err
	}
	var fixedTag string
	if subproject == "kubernetes" {
		fixedTag = fmt.Sprintf("v%s", tag)
	} else {
		fixedTag = fmt.Sprintf("kubernetes-%s", tag)
	}
	return checkoutTree(repo, fixedTag)
}

func retrieveOrCloneKubernetesSubprojectRepo(subproject string) (*git.Repository, error) {
	needsClone := false
	if cmdOptions.kubernetesPath == "" {
		needsClone = true
	} else {
		_, err := os.Stat(path.Join(cmdOptions.kubernetesPath, subproject))
		needsClone = os.IsNotExist(err)
	}
	projectPath := path.Join(cmdOptions.kubernetesPath, subproject)
	if needsClone {
		projectUrl := fmt.Sprintf("https://github.com/kubernetes/%s", subproject)
		if cmdOptions.kubernetesPath == "" {
			fmt.Fprintf(os.Stderr, "cloning project %s in memory\n", projectUrl)
		} else {
			fmt.Fprintf(os.Stderr, "project %s not found; cloning project %s in memory\n", projectPath, projectUrl)
		}
		return git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL: projectUrl,
		})
	}
	return git.PlainOpen(projectPath)
}

var rootCmd = &cobra.Command{
	Use: "vendor-kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		printGoMod()
	},
}

func printGoMod() {
	_, tree, err := retrieveOrCloneKubernetesSubproject("kubernetes", cmdOptions.kubernetesTag)
	if err != nil {
		log.Fatalf("could not retrieve kubernetes repo: %v", err)
	}
	subprojects, err := kubernetesSubprojects(tree)
	if err != nil {
		log.Fatalf("could not determine kubernetes subprojects: %v", err)
	}
	replaceList := []string{}
	for _, subproject := range subprojects {
		subprojectCommit, _, err := retrieveOrCloneKubernetesSubproject(subproject, cmdOptions.kubernetesTag)
		if err != nil {
			log.Fatalf("could not clone subproject %s: %v", subproject, err)
		}
		year, month, day := subprojectCommit.Committer.When.Date()
		hour, minute, second := subprojectCommit.Committer.When.Clock()
		date := fmt.Sprintf("%d%.2d%.2d%.2d%.2d%.2d", year, month, day, hour, minute, second)
		revision := subprojectCommit.Hash
		replaceList = append(replaceList, fmt.Sprintf("k8s.io/%s => k8s.io/%s v0.0.0-%s-%.12s", subproject, subproject, date, revision))
	}

	fmt.Println("require (")
	fmt.Printf("  k8s.io/kubernetes v%s\n", cmdOptions.kubernetesTag)
	fmt.Println(")")
	fmt.Println()
	fmt.Println("replace (")
	for _, replace := range replaceList {
		fmt.Printf("  %s\n", replace)
	}
	fmt.Println(")")
}

func main() {
	rootCmd.PersistentFlags().StringVar(&cmdOptions.kubernetesPath, "kubernetes-path", "", "(optional) Path pointing to the Kubernetes repository. (e.g. \"~/projects/go/src/k8s.io\"). If not provided or if some repositories are missing, they will be cloned in memory")
	rootCmd.PersistentFlags().StringVar(&cmdOptions.kubernetesTag, "kubernetes-tag", "", "Kubernetes tag to build go.mod for (e.g. \"1.15.3\")")
	rootCmd.MarkPersistentFlagRequired("kubernetes-tag")
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
