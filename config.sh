eval $(minikube -p minikube docker-env)
export POSTGRES_DB=$(kubectl get configmap --namespace default postgres-configuration -o jsonpath="{.data.POSTGRES_DB}" )
export POSTGRES_USER=$(kubectl get configmap --namespace default postgres-configuration -o jsonpath="{.data.POSTGRES_USER}" )
export DBPSW=$(kubectl get configmap --namespace default postgres-configuration -o jsonpath="{.data.POSTGRES_PASSWORD}" )
export DBHOST=$(kubectl get service postgres-service -o jsonpath="{.spec.clusterIP}")
docker build . -t golangdb --build-arg DBPSW --build-arg DBHOST --build-arg POSTGRES_DB --build-arg POSTGRES_USER 
# kubectl apply -f dbclient.yml
# or w/o kubernetes only docker
# docker run -e DBHOST -e DBPSW -it golangdb

DB_POD_NAME=$(kubectl get pods | grep golang-db | cut -d ' ' -f1)
echo $DB_POD_NAME
kubectl port-forward $DB_POD_NAME 8000:8080
echo $POSTGRES_DB $POSTGRES_USER $DBPSW $DBHOST
kubectl exec -ti postgres-statefulset-0 -- psql -U $POSTGRES_USER -d $POSTGRES_DB