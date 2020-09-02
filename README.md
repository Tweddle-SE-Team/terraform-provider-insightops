InsightOps Terraform Provider
=============================

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

###### Powered by: https://www.terraform.io

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.13.x
-	[Go](https://golang.org/doc/install) 1.14 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/Tweddle-SE-Team/terraform-provider-insightops`

```sh
$ go get github.com/Tweddle-SE-Team/terraform-provider-insightops
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/Tweddle-SE-Team/terraform-provider-insightops
$ go build
```

Using the provider
------------------

Refer to the READMEs inside the [examples](https://github.com/Tweddle-SE-Team/terraform-provider-insightops/examples) folder to
see how to configure each resource provided by this terraform provider.

Developing the Provider
-----------------------

To compile the provider, run `go build`. This will build the provider and put the provider binary in the current
`$GOPATH/src/github.com/Tweddle-SE-Team/terraform-provider-insightops` directory.

```sh
$ go build
...
$ ls -la terraform-provider-insightops
...
```

In order to test the provider, you can simply run:

```sh
$ TF_ACC=1 INSIGHTOPS_API_KEY="API_KEY" go test $(go list ./...) -timeout 120m -v
```
Expected output:

```
$ TF_ACC=1 INSIGHTOPS_API_KEY="<API_KEY>" go test $(go list ./...) -timeout 120m -v
?       github.com/Tweddle-SE-Team/terraform-provider-insightops [no test files]
=== RUN   TestInsightOpsProvider
--- PASS: TestInsightOpsProvider (0.00s)
=== RUN   TestAccInsightOpsLog_Create
--- PASS: TestAccInsightOpsLog_Create (8.81s)
=== RUN   TestAccInsightOpsLog_Update
--- PASS: TestAccInsightOpsLog_Update (11.08s)
=== RUN   TestAccInsightOpsLogSets_Create
--- PASS: TestAccInsightOpsLogSets_Create (0.98s)
=== RUN   TestAccInsightOpsLogSets_Update
--- PASS: TestAccInsightOpsLogSets_Update (1.60s)
=== RUN   TestAccInsightOpsTags_Create
--- PASS: TestAccInsightOpsTags_Create (13.36s)
=== RUN   TestAccInsightOpsTags_Update
--- PASS: TestAccInsightOpsTags_Update (19.71s)
PASS
ok      github.com/Tweddle-SE-Team/terraform-provider-insightops/insightops      55.636s

```

Or specific tests can also be executed as follows:

```sh
$ TF_ACC=1 INSIGHTOPS_API_KEY="<API_KEY>" INSIGHTOPS_REGION="<REGION>" go test github.com/Tweddle-SE-Team/terraform-provider-insightops/insightops -run  ^TestAccInsightOpsTags_Create$ -timeout 120m -v
```

The acceptance tests require a INSIGHTOPS_API_KEY and INSIGHTOPS_REGION to be set. These env variables value will be used within the tests to
successfully interact with the InsightOps api.

*Note: Acceptance tests create real resources and perform clean up tasks afterwards.*

Contributing
------------
Please follow the guidelines from:

 - [Contributor Guidelines](.github/CONTRIBUTING.md)

Authors
-------

Daniel I. Khan Ramiro

See also the list of [contributors](https://github.com/Tweddle-SE-Team/terraform-provider-insightops/graphs/contributors) who
participated in this project.
