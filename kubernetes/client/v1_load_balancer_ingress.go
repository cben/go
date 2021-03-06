/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.
type V1LoadBalancerIngress struct {

	// Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)
	Hostname string `json:"hostname,omitempty"`

	// IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)
	Ip string `json:"ip,omitempty"`
}
