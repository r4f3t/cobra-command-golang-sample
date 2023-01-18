
## Cobra installation and Usage

Installing cobra package;

```bash
  go get -u github.com/spf13/cobra@latest
```

Installing cli tool;

 ```bash
  go install github.com/spf13/cobra@latest
``` 

After installing tool you can use this command for creation root structure;
 ```bash
    cobra-cli init
``` 

After this command you will have structure like below;

 ```bash
├── LICENSE
├── cmd
│   └── root.go
├── go.mod
├── go.sum
└── main.go
``` 

Then you can create your project/go files by using ;
 ```bash
    cobra-cli add api
``` 

 ```bash
    cobra-cli add job
``` 

After that you will have structure like below;

 ```bash
├── LICENSE
├── cmd
│   └── root.go
│   └── api.go
│   └── job.go
├── go.mod
├── go.sum
└── main.go
``` 

Now you can manage your projects inside of related go file in cmd.
For 'api' in this repo I used echo for create a web api.
You can see initialization of it inside of cmd/api.go .
