// Copyright 2019 Adam Shannon
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package cidetails

import (
	"os"
	"strings"
)

type vendor struct {
	// case-insensitive names to describe a CI service
	// names should be from the current to oldest
	names  []string
	envVar string
	pr     func() bool
}

func n(names ...string) []string {
	return names
}

func nonempty(envVar string) func() bool {
	return func() bool {
		return strings.TrimSpace(os.Getenv(envVar)) != ""
	}
}

// func any(envVars ...string) bool { // TODO(adam): which CI provider was this for?
// 	for i := range envVars {
// 		if nonempty(envVars[i])() {
// 			return true
// 		}
// 	}
// 	return false
// }

var vendors = []vendor{
	{
		names:  n("AppVeyor"),
		envVar: "APPVEYOR",
		pr:     nonempty("APPVEYOR_PULL_REQUEST_NUMBER"),
	},
	{
		names:  n("AWS CodeBuild"),
		envVar: "CODEBUILD_BUILD_ARN",
		pr:     func() bool { return false },
	},
	{
		names:  n("Azure Pipelines"),
		envVar: "SYSTEM_TEAMFOUNDATIONCOLLECTIONURI",
		pr:     nonempty("SYSTEM_PULLREQUEST_PULLREQUESTID"),
	},
	{
		names:  n("Bamboo"),
		envVar: "bamboo_planKey",
		pr:     func() bool { return false },
	},
	{
		names:  n("Bitbucket Pipelines"),
		envVar: "BITBUCKET_COMMIT",
		pr:     nonempty("BITBUCKET_PR_ID"),
	},
	{
		names:  n("Bitrise"),
		envVar: "BITRISE_IO",
		pr:     nonempty("BITRISE_PULL_REQUEST"),
	},
	{
		names:  n("Buddy"),
		envVar: "BUDDY_WORKSPACE_ID",
		pr:     nonempty("BUDDY_EXECUTION_PULL_REQUEST_ID"),
	},
	{
		names:  n("Buildkite"),
		envVar: "BUILDKITE",
		pr:     func() bool { return os.Getenv("BUILDKITE_PULL_REQUEST") != "false" },
	},
	{
		names:  n("CircleCI"),
		envVar: "CIRCLECI",
		pr:     nonempty("CIRCLE_PULL_REQUEST"),
	},
	{
		names:  n("Cirrus CI"),
		envVar: "CIRRUS_CI",
		pr:     nonempty("CIRRUS_PR"),
	},
	{
		names:  n("Codeship"),
		envVar: "CI_NAME",
		pr:     func() bool { return false },
	},
	{
		names:  n("Drone"),
		envVar: "DRONE",
		pr:     nonempty("DRONE_BUILD_EVENT"),
	},
	{
		names:  n("dsari"),
		envVar: "DSARI",
		pr:     func() bool { return false },
	},
	{
		names:  n("GitLab CI"),
		envVar: "GITLAB_CI",
		pr:     func() bool { return false },
	},
	{
		names:  n("GoCD"),
		envVar: "GO_PIPELINE_LABEL",
		pr:     func() bool { return false },
	},
	{
		names:  n("Hudson"),
		envVar: "HUDSON_URL",
		pr:     func() bool { return false },
	},
	{
		names:  n("Jenkins"),
		envVar: "JENKINS_URL",
		pr:     nonempty("ghprbPullId"),
	},
	{
		names:  n("Magnum CI"),
		envVar: "MAGNUM",
		pr:     func() bool { return false },
	},
	{
		names:  n("Netlify CI"),
		envVar: "NETLIFY_BUILD_BASE",
		pr:     func() bool { return os.Getenv("PULL_REQUEST") != "false" },
	},
	{
		names:  n("Nevercode"),
		envVar: "NEVERCODE",
		pr:     func() bool { return os.Getenv("NEVERCODE_PULL_REQUEST") != "false" },
	},
	{
		names:  n("Sail CI"),
		envVar: "SAILCI",
		pr:     nonempty("SAIL_PULL_REQUEST_NUMBER"),
	},
	{
		names:  n("Semaphore"),
		envVar: "SEMAPHORE",
		pr:     nonempty("PULL_REQUEST_NUMBER"),
	},
	{
		names:  n("Shippable"),
		envVar: "SHIPPABLE",
		pr:     func() bool { return os.Getenv("IS_PULL_REQUEST") == "true" },
	},
	{
		names:  n("Solano CI"),
		envVar: "TDDIUM",
		pr:     nonempty("TDDIUM_PR_ID"),
	},
	{
		names:  n("Strider CD"),
		envVar: "STRIDER",
		pr:     func() bool { return false },
	},
	{
		names:  n("TeamCity"),
		envVar: "TEAMCITY_VERSION",
		pr:     func() bool { return false },
	},
	{
		names:  n("Travis CI", "TravisCI"),
		envVar: "TRAVIS",
		pr:     func() bool { return os.Getenv("TRAVIS_PULL_REQUEST") != "false" },
	},
}
