/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type CoreTaskMetadata struct {
	// Indicates whether the system should attempt to lookup this task's output to avoid duplication of work.
	Discoverable bool `json:"discoverable,omitempty"`
	// Runtime information about the task.
	Runtime *CoreRuntimeMetadata `json:"runtime,omitempty"`
	// The overall timeout of a task including user-triggered retries.
	Timeout string `json:"timeout,omitempty"`
	// Number of retries per task.
	Retries *CoreRetryStrategy `json:"retries,omitempty"`
	// Indicates a logical version to apply to this task for the purpose of discovery.
	DiscoveryVersion string `json:"discovery_version,omitempty"`
	// If set, this indicates that this task is deprecated.  This will enable owners of tasks to notify consumers of the ending of support for a given task.
	DeprecatedErrorMessage string `json:"deprecated_error_message,omitempty"`
	Interruptible bool `json:"interruptible,omitempty"`
	CacheSerializable bool `json:"cache_serializable,omitempty"`
	// Indicates whether the task will generate a Deck URI when it finishes executing.
	GeneratesDeck bool `json:"generates_deck,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	// pod_template_name is the unique name of a PodTemplate k8s resource to be used as the base configuration if this task creates a k8s Pod. If this value is set, the specified PodTemplate will be used instead of, but applied identically as, the default PodTemplate configured in FlytePropeller.
	PodTemplateName string `json:"pod_template_name,omitempty"`
}
