// Define the function newServiceForPod which takes a pointer to a RoarApp custom resource (CR) and returns a pointer to a Service object.
func newServiceForPod(cr *roarappv1alpha1.RoarApp) *corev1.Service {

    // Convert the global variable nextPort (an integer) to a string using strconv.Itoa.
    // This is necessary because we will use this port number as part of the service's name to ensure uniqueness.
    strPort := strconv.Itoa(nextPort)

    // Define a map of string keys to string values to hold labels for the service.
    // Here, we create a label with the key "app" and set its value to the name of the RoarApp CR.
    // These labels are used by the service to select which pods to target.
    labels := map[string]string{
        "app": cr.Name,
    }

    // Return a pointer to a new Service object configured for the RoarApp.
    return &corev1.Service{
        // ObjectMeta contains metadata about the service, like its name and namespace.
        ObjectMeta: metav1.ObjectMeta{
            // Construct the service name by concatenating the RoarApp's name, a "-service-" string, and the strPort.
            // This ensures that each service has a unique name based on the RoarApp's name and the next available port.
            Name:      cr.Name + "-service-" + strPort,
            // Set the namespace of the service to be the same as the RoarApp's namespace.
            // This ensures that the service is created in the correct Kubernetes namespace.
            Namespace: cr.Namespace,
        },
        // Spec defines the behavior of the service.
        Spec: corev1.ServiceSpec{
            // Selector uses the labels map defined earlier to select the pods that this service will target.
            // This means the service will route traffic to pods that have the label "app" with the value of the RoarApp's name.
            Selector: labels,
            // Ports is an array of ServicePort objects. Here, we define a single port that the service will expose.
            Ports: []corev1.ServicePort{{
                // Protocol specifies the protocol used by the service, in this case, TCP.
                Protocol:   corev1.ProtocolTCP,
                // Port specifies the port number that the service will expose externally.
                Port:       8089,
                // TargetPort specifies the port on the pods that the service will forward traffic to.
                // Here, it forwards to port 8080 on the pods.
                TargetPort: intstr.FromInt(8080),
                // NodePort specifies the port number on each node's IP where this service will be exposed.
                // Here, it uses the global variable nextPort, allowing external access to the service.
                NodePort:   int32(nextPort),
            }},
            // Type specifies the type of service. NodePort means the service will be accessible outside the cluster
            // via an assigned port on each node's IP.
            Type: corev1.ServiceTypeNodePort,
        },
    }
}
