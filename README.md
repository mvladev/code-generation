Code generation bug

```bash
$ ./vendor/k8s.io/code-generator/generate-groups.sh \
  all \
  github.com/mvladev/code-generation/client \
  github.com/mvladev/code-generation/apis \
  "example:v1" \
  -h dummy


Generating deepcopy funcs
Generating clientset for example:v1 at github.com/mvladev/code-generation/client/clientset
ERROR: logging before flag.Parse: F1106 07:49:42.418067   80131 main.go:69] Error: Failed executing generator: some packages had errors:
errors in package "github.com/mvladev/code-generation/client/clientset/versioned":
unable to format file "/Users/i068969/git/go/src/github.com/mvladev/code-generation/client/clientset/versioned/clientset.go" (6:6: expected 'STRING', found '-' (and 10 more errors)).

errors in package "github.com/mvladev/code-generation/client/clientset/versioned/scheme":
unable to format file "/Users/i068969/git/go/src/github.com/mvladev/code-generation/client/clientset/versioned/scheme/register.go" (8:6: expected 'STRING', found '-' (and 4 more errors)).

errors in package "github.com/mvladev/code-generation/client/clientset/versioned/fake":
unable to format file "/Users/i068969/git/go/src/github.com/mvladev/code-generation/client/clientset/versioned/fake/clientset_generated.go" (4:10: expected 'STRING', found '-' (and 9 more errors)).
unable to format file "/Users/i068969/git/go/src/github.com/mvladev/code-generation/client/clientset/versioned/fake/register.go" (4:6: expected 'STRING', found '-' (and 4 more errors)).

errors in package "github.com/mvladev/code-generation/client/clientset/versioned/typed/code-generator/v1":
unable to format file "/Users/i068969/git/go/src/github.com/mvladev/code-generation/client/clientset/versioned/typed/code-generator/v1/code-generator_client.go" (11:10: expected type, found '-' (and 2 mor
e errors)).
unable to format file "/Users/i068969/git/go/src/github.com/mvladev/code-generation/client/clientset/versioned/typed/code-generator/v1/testtype.go" (40:26: missing ',' in parameter list).

errors in package "github.com/mvladev/code-generation/client/clientset/versioned/typed/code-generator/v1/fake":
unable to format file "/Users/i068969/git/go/src/github.com/mvladev/code-generation/client/clientset/versioned/typed/code-generator/v1/fake/fake_testtype.go" (16:16: expected ';', found '-' (and 1 more er
rors)).
unable to format file "/Users/i068969/git/go/src/github.com/mvladev/code-generation/client/clientset/versioned/typed/code-generator/v1/fake/fake_code-generator_client.go" (10:14: expected type, found '-'
(and 2 more errors)).
```
