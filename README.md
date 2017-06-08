# k8aos

```                                                                  


 ____    ____         ____           _____             ______  
|    |  |    |   ____|\   \     ____|\    \        ___|\     \
|    |  |    |  /    /\    \   /     /\    \      |    |\     \
|    | /    // |    |  |    | /     /  \    \     |    |/____/|
|    |/ _ _//  |    |__|    ||     |    |    | ___|    \|   | |
|    |\    \'  |    .--.    ||     |    |    ||    \    \___|/
|    | \    \  |    |  |    ||\     \  /    /||    |\     \    
|____|  \____\ |____|  |____|| \_____\/____/ ||\ ___\|_____|   
|    |   |    ||    |  |    | \ |    ||    | /| |    |     |   
|____|   |____||____|  |____|  \|____||____|/  \|____|_____|   
  \(       )/    \(      )/       \(    )/        \(    )/     
   '       '      '      '         '    '          '    '      

```

This tool is a chaos-monkey for deleting kubernetes containers to test system stability and healing.

## Installation
```
go get github.com/AlexsJones/k8aos
```

## Configuration

`kubectl config view > config`

You may want to override the URL and use `kubectl proxy` with `http://localhost:8001` in the config file
to avoid having to use CAFiles or TLS issues.

## Usage

```
Commands:
  again         Run the last mischief command again
  clear         clear the screen
  connect       Provide an absolute path to config as second argument e.g. connect /tmp/config
  exit          exit the program
  help          display help
  inspect       inspect the current cluster containers
  mischief      Destroy a pod in a random namespace (can specify with second argument)
```

Essentially you connect to a cluster and k8aos will control deletion of pods and timing of events for you.
