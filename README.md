# web-golang
web in go language from wiki example

- CI deployment to GKE via github actions

- end to end testing with cypress (screenshot of dashboard run below)

![immagine](https://user-images.githubusercontent.com/3272563/157900395-1fe1799a-1628-43c3-925a-863a15d53860.png)

## docker or minikube

Without a google cloud subscription, do for example
```
sudo systemctl start docker
```
Either you simply
```
docker run -p 80:8080 giuliohome/web.golang:latest
```
and visit http://localhost/view/a1

or you can
```
minikube start
```
then deploy the demo to the local node
```
minikube kubectl -- apply -f deployment.yml
```
to get an external ip
```
minikube tunnel -c --log_dir tunnel_log/
minikube kubectl -- get all
```
Open the browser and check it out

Finally reclaim all the sapce
```
minikube stop
minikube delete --purge --all
```
and eventually
```
docker system prune -a
```

## end to end testing

useful commands, especially for local dev
```
npx cypress run --spec cypress\integration\1-getting-started\web-golang.js
```
and 
```
npx cypress open
```
![image](https://user-images.githubusercontent.com/3272563/158066376-db8e0e1f-7442-4bda-bc17-cd75b1979529.png)

## lens + minikube + stateful db pod 

log from golang console

![image](https://user-images.githubusercontent.com/3272563/158406574-13008fcd-6c6e-48e4-b81a-25bfd4b72caa.png)





