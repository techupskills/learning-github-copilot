// newServiceForPod is a function that creates a new Kubernetes Service for a given RoarApp instance.
// This service is designed to expose the RoarApp's pods on a specific node port, allowing external access.
//
// Parameters:
// - cr: A pointer to the RoarApp custom resource instance. This parameter is crucial as it contains
//       the metadata and specifications of the RoarApp instance, such as its name and namespace,
//       which are used to configure the service accordingly.
//
// Returns:
// - A pointer to a corev1.Service object. This object represents the newly created service configured
//   to expose the RoarApp instance. The service is not yet applied to the Kubernetes cluster; it merely
//   represents the desired state.
//
// Detailed Steps:
// 1. Convert the global variable `nextPort` to a string using strconv.Itoa. This port number is used
//    to uniquely identify the service and is part of the service's name to ensure uniqueness across
//    all services created by this function.
// 2. Define a set of labels in a map[string]string. These labels are applied to the service and are
//    used to select the pods that the service will expose. In this case, the label is based on the
//    RoarApp's name, ensuring that the service only selects pods belonging to the specific RoarApp instance.
// 3. Construct a corev1.Service object. This object includes several important fields:
//    - ObjectMeta: Contains metadata about the service, such as its name and namespace. The name is
//                  constructed by combining the RoarApp's name with "-service-" and the string representation
//                  of `nextPort`, ensuring uniqueness. The namespace is the same as that of the RoarApp instance,
//                  ensuring that the service is created in the correct context.
//    - Spec: Defines the specifications of the service. This includes:
//            - Selector: A map that specifies which pods the service should expose. It uses the labels
//                        defined earlier to select the appropriate pods.
//            - Ports: A list of service ports. Each entry defines a port that the service will expose.
//                     This includes the protocol (TCP), the port number on the service (8089), the target
//                     port on the pods (8080), and the node port (nextPort). The node port is what allows
//                     external access to the service.
//            - Type: Specifies the type of service. NodePort is used here, which means the service will
//                    be accessible through an external port on each node.
//
// Note: The `nextPort` variable must be managed externally to ensure it is unique and within the valid
//       range for node ports in the Kubernetes cluster. This function assumes `nextPort` has been
//       appropriately incremented and managed to avoid conflicts with other services.

