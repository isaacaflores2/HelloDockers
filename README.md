# HelloDockers
Collection of docker images in various languages that say Hello
###### Current languages: Go, C#
---

### To create the shell image run the below commands:
```
docker build -t isaacaflores/hello .
docker run --rm isaacaflores/hello
```

### To add your image to your repository run:
```
docker login
docker push isaacaflores/hello
```

![dockerlogo](docker.png)
