# CGO_ENABLED=0 go build db.go
FROM busybox
COPY ./db /db
ARG DBPSW
ENV DBPSW=$DBPSW
ARG DBHOST
ENV DBHOST=$DBHOST
ARG POSTGRES_USER
ENV POSTGRES_USER=$POSTGRES_USER
ARG POSTGRES_DB
ENV POSTGRES_DB=$POSTGRES_DB
CMD ["/db"]
# eval $(minikube -p minikube docker-env)
# export POSTGRES_DB=$(kubectl get configmap --namespace default postgres-configuration -o jsonpath="{.data.POSTGRES_DB}" )
# export POSTGRES_USER=$(kubectl get configmap --namespace default postgres-configuration -o jsonpath="{.data.POSTGRES_USER}" )
# export DBPSW=$(kubectl get configmap --namespace default postgres-configuration -o jsonpath="{.data.POSTGRES_PASSWORD}" )
# export DBHOST=$(kubectl get service postgres-service -o jsonpath="{.spec.clusterIP}")
# docker build . -t golangdb --build-arg DBPSW --build-arg DBHOST --build-arg POSTGRES_DB --build-arg POSTGRES_USER 
# kubectl apply -f dbclient.yml
# or w/o kubernetes only docker
# docker run -e DBHOST -e DBPSW -it golangdb
