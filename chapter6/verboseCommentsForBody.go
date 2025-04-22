// newServiceForPod creates a new Kubernetes Service resource for a given RoarApp custom resource.
// The Service is configured as a NodePort service, which exposes a specific port for external access.
// This function ensures that the Service is uniquely named and associated with the corresponding RoarApp instance.
//
// Parameters:
//   - cr: A pointer to a RoarApp custom resource instance.
//
// Returns:
//   - A pointer to a corev1.Service object configured with the specified labels, ports, and namespace.
func newServiceForPod(cr *roarappv1alpha1.RoarApp) *corev1.Service {

    // Convert the next available port (global variable `nextPort`) to a string.
    // This port will be used to uniquely identify the Service and its associated Pod.
    // The `strconv.Itoa` function converts the integer `nextPort` to its string representation.
    strPort := strconv.Itoa(nextPort)

    // Define a set of labels to associate the Service with the corresponding RoarApp instance.
    // These labels will be used by the Service to select the appropriate Pods.
    // The "app" label is set to the name of the RoarApp instance, ensuring that the Service
    // can target Pods created for this specific RoarApp instance.
    labels := map[string]string{
        "app": cr.Name, // The "app" label is set to the name of the RoarApp instance.
    }

    // Create and return a new Service object.
    // The Service is configured with metadata and specifications to expose the application.
    return &corev1.Service{
        ObjectMeta: metav1.ObjectMeta{
            // The name of the Service is constructed using the RoarApp name and the port.
            // This ensures that each Service has a unique name, even if multiple Services
            // are created for the same RoarApp instance.
            Name: cr.Name + "-service-" + strPort,

            // The Service is created in the same namespace as the RoarApp instance.
            // This ensures that the Service is scoped to the same namespace as its associated Pods.
            Namespace: cr.Namespace,
        },
        Spec: corev1.ServiceSpec{
            // The selector specifies which Pods this Service will target.
            // It uses the labels defined earlier to match the Pods created for this RoarApp instance.
            Selector: labels,

            // Define the ports that the Service will expose.
            Ports: []corev1.ServicePort{{
                // The protocol used by the Service is TCP, which is the default protocol for most applications.
                Protocol: corev1.ProtocolTCP,

                // The port exposed by the Service for external access.
                // This is the port that clients outside the cluster will use to access the application.
                Port: 8089,

                // The port on the target Pods that the Service will forward traffic to.
                // This is the port where the application inside the Pods is listening for incoming traffic.
                TargetPort: intstr.FromInt(8080),

                // The NodePort assigned to this Service for external access.
                // This is a unique port on each node in the cluster that forwards traffic to the Service.
                NodePort: int32(nextPort),
            }},

            // The Service is of type NodePort, which makes it accessible from outside the cluster.
            // NodePort Services expose the application on a static port on each node's IP address.
            Type: corev1.ServiceTypeNodePort,
        },
    }
}
