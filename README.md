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


## Installation
```
go get github.com/AlexsJones/k8aos
```

## Configuration

`kubectl config view > config`

You may want to override the URL and use `kubectl proxy` with `http://localhost:8001` in the config file
to avoid having to use CAFiles or TLS issues.
