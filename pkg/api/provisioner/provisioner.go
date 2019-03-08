package provisioner

import (
	"github.com/yard-turkey/lib-bucket-provisioner/pkg/apis/objectbucket.io/v1alpha1"
)

// Provisioner the interface to be implemented by users of this
// library and executed by the Reconciler
type Provisioner interface {
	// Provision should be implemented to handle bucket creation
	// for the target object store
	Provision(options *BucketOptions) (*v1alpha1.ObjectBucket, error)
	// Delete should be implemented to handle bucket deletion
	// for the target object store
	Delete(ob *v1alpha1.ObjectBucket) error
}

type BucketOptions struct {
	// ReclaimPolicy is the reclaimPolicy of the OBC's storage class
	ReclaimPolicy v1alpha1.ReclaimPolicy
	// ObjectBucketName is the name of the ObjectBucket API resource
	ObjectBucketName string
	// BucketNae is the name of the bucket within the object store
	BucketName string
	// ObjectBucketClaim is a pointer to the initiating OBC object
	ObjectBucketClaim *v1alpha1.ObjectBucketClaim
	// Parameters is a complete copy of the OBC's storage class Parameters field
	Parameters map[string]string
}
