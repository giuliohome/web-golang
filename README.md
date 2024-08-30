
# DigitalOcean setup instructions



 - Create cluster 
   - at least 2 nodes x (2GB RAM, 1 vCPU, 50GB storage)
   - 2 x $12/month


 - Container Registry =
   - Integrating with Kubernetes > Save

  - GitHub 
    - update action secret CLUSTER_NAME

 - Kubernetes > Marketplace 
   - Install Nginx Ingress Controller

 - Networking 
   - IP from Load Balancer
   - create new A record in DNS Domain for http host
   - 30 TTL (seconds)

 - Kubernetes > Marketplace 
   - Install Cert-Manager

 - Github
    - Run Workflow

Optional, for troubleshooting
- Gitpod
  - install doctl
  - install kubectl