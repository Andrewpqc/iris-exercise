cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--kubeconfig=/etc/kubernetes/kubelet.conf --require-kubeconfig=true"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=systemd"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_EXTRA_ARGS


[Service]
Environment="KUBELET_KUBECONFIG_ARGS=--kubeconfig=/etc/kubernetes/kubelet.conf --require-kubeconfig=true"
Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"
Environment="KUBELET_NETWORK_ARGS=--network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/opt/cni/bin"
Environment="KUBELET_DNS_ARGS=--cluster-dns=10.96.0.10 --cluster-domain=cluster.local"
Environment="KUBELET_AUTHZ_ARGS=--authorization-mode=Webhook --client-ca-file=/etc/kubernetes/pki/ca.crt"
Environment="KUBELET_CADVISOR_ARGS=--cadvisor-port=0"
Environment="KUBELET_CGROUP_ARGS=--cgroup-driver=cgroupfs"
ExecStart=
ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_EXTRA_ARGS



systemctl status kubelet
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; enabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf, 20-pod-infra-image.conf
   Active: activating (auto-restart) (Result: exit-code) since 日 2018-04-22 10:26:38 CST; 1s ago
     Docs: http://kubernetes.io/docs/
  Process: 21169 ExecStart=/usr/bin/kubelet $KUBELET_KUBECONFIG_ARGS $KUBELET_SYSTEM_PODS_ARGS $KUBELET_NETWORK_ARGS $KUBELET_DNS_ARGS $KUBELET_AUTHZ_ARGS $KUBELET_CADVISOR_ARGS $KUBELET_CGROUP_ARGS $KUBELET_EXTRA_ARGS (code=exited, status=255)
 Main PID: 21169 (code=exited, status=255)

4月 22 10:26:38 iZwz9f6pgul78p7die5tlzZ kubelet[21169]: --seccomp-profile-root string                         <Warning: Alpha feature> ...comp")
4月 22 10:26:38 iZwz9f6pgul78p7die5tlzZ kubelet[21169]: --stderrthreshold severity                            logs at or above this thr...ult 2)
4月 22 10:26:38 iZwz9f6pgul78p7die5tlzZ kubelet[21169]: -v, --v Level                                             log level for V logs
4月 22 10:26:38 iZwz9f6pgul78p7die5tlzZ kubelet[21169]: --version version[=true]                              Print version information and quit
4月 22 10:26:38 iZwz9f6pgul78p7die5tlzZ kubelet[21169]: --vmodule moduleSpec                                  comma-separated list of p...ogging
4月 22 10:26:38 iZwz9f6pgul78p7die5tlzZ systemd[1]: kubelet.service: main process exited, code=exited, status=255/n/a
4月 22 10:26:38 iZwz9f6pgul78p7die5tlzZ kubelet[21169]: --volume-plugin-dir string                            The full path of the dire...xec/")
4月 22 10:26:38 iZwz9f6pgul78p7die5tlzZ kubelet[21169]: F0422 10:26:38.470722   21169 server.go:145] unknown flag: --require-kubeconfig
4月 22 10:26:38 iZwz9f6pgul78p7die5tlzZ systemd[1]: Unit kubelet.service entered failed state.
4月 22 10:26:38 iZwz9f6pgul78p7die5tlzZ systemd[1]: kubelet.service failed.


[root@iZwz9f6pgul78p7die5tlzZ kubelet.service.d]# tail -n 10 /var/log/messages 
Apr 22 10:46:38 iZwz9f6pgul78p7die5tlzZ kubelet: Flag --authorization-mode has been deprecated, This parameter should be set via the config file specified by the Kubelet's --config flag. See https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/ for more information.
Apr 22 10:46:38 iZwz9f6pgul78p7die5tlzZ kubelet: Flag --client-ca-file has been deprecated, This parameter should be set via the config file specified by the Kubelet's --config flag. See https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/ for more information.
Apr 22 10:46:38 iZwz9f6pgul78p7die5tlzZ kubelet: Flag --cadvisor-port has been deprecated, The default will change to 0 (disabled) in 1.12, and the cadvisor port will be removed entirely in 1.13
Apr 22 10:46:38 iZwz9f6pgul78p7die5tlzZ kubelet: Flag --cgroup-driver has been deprecated, This parameter should be set via the config file specified by the Kubelet's --config flag. See https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/ for more information.
Apr 22 10:46:38 iZwz9f6pgul78p7die5tlzZ kubelet: Flag --fail-swap-on has been deprecated, This parameter should be set via the config file specified by the Kubelet's --config flag. See https://kubernetes.io/docs/tasks/administer-cluster/kubelet-config-file/ for more information.
Apr 22 10:46:38 iZwz9f6pgul78p7die5tlzZ kubelet: I0422 10:46:38.716591   21994 feature_gate.go:226] feature gates: &{{} map[]}
Apr 22 10:46:38 iZwz9f6pgul78p7die5tlzZ kubelet: F0422 10:46:38.716693   21994 server.go:218] unable to load client CA file /etc/kubernetes/pki/ca.crt: open /etc/kubernetes/pki/ca.crt: no such file or directory
Apr 22 10:46:38 iZwz9f6pgul78p7die5tlzZ systemd: kubelet.service: main process exited, code=exited, status=255/n/a
Apr 22 10:46:38 iZwz9f6pgul78p7die5tlzZ systemd: Unit kubelet.service entered failed state.
Apr 22 10:46:38 iZwz9f6pgul78p7die5tlzZ systemd: kubelet.service failed.

taints




[root@iZwz9f6pgul78p7die5tlzZ ~]# kubeadm init --apiserver-advertise-address=39.108.79.110 --kubernetes-version=v1.7.5 --pod-network-cidr=10.244.0.0/12 --skip-preflight-checks
[kubeadm] WARNING: kubeadm is in beta, please do not use it for production clusters.
[init] Using Kubernetes version: v1.7.5
[init] Using Authorization modes: [Node RBAC]
[preflight] Skipping pre-flight checks
[kubeadm] WARNING: starting in 1.8, tokens expire after 24 hours by default (if you require a non-expiring token use --token-ttl 0)
[certificates] Using the existing CA certificate and key.
[certificates] Using the existing API Server certificate and key.
[certificates] Using the existing API Server kubelet client certificate and key.
[certificates] Using the existing service account token signing key.
[certificates] Using the existing front-proxy CA certificate and key.
[certificates] Using the existing front-proxy client certificate and key.
[certificates] Valid certificates and keys now exist in "/etc/kubernetes/pki"
[kubeconfig] Using existing up-to-date KubeConfig file: "/etc/kubernetes/admin.conf"
[kubeconfig] Using existing up-to-date KubeConfig file: "/etc/kubernetes/kubelet.conf"
[kubeconfig] Using existing up-to-date KubeConfig file: "/etc/kubernetes/controller-manager.conf"
[kubeconfig] Using existing up-to-date KubeConfig file: "/etc/kubernetes/scheduler.conf"
[apiclient] Created API client, waiting for the control plane to become ready
[apiclient] All control plane components are healthy after 40.004549 seconds
[token] Using token: 270890.82636121ccc09874
[apiconfig] Created RBAC rules
[addons] Applied essential addon: kube-proxy
[addons] Applied essential addon: kube-dns

Your Kubernetes master has initialized successfully!

To start using your cluster, you need to run (as a regular user):

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  http://kubernetes.io/docs/admin/addons/

You can now join any number of machines by running the following on each node
as root:

  kubeadm join --token 270890.82636121ccc09874 39.108.79.110:6443

[root@iZwz9f6pgul78p7die5tlzZ ~]# 


sudo systemctl daemon-reload && sudo systemctl restart kubelet docker
export kubever=$(kubectl version | base64 | tr -d '\n')
kubectl delete -f "https://cloud.weave.works/k8s/net?k8s-version=$kubever"
kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$kubever"
