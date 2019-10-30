package api

import (
	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	"github.com/oasislabs/oasis-core/go/common/sgx"
)

// PolicySGXSignatureContext is the context used to sign PolicySGX documents.
var PolicySGXSignatureContext = signature.NewContext("EkKmPolS")

// PolicySGX is a key manager access control policy for the replicated
// SGX key manager.
type PolicySGX struct {
	// Serial is the monotonically increasing policy serial number.
	Serial uint32 `json:"serial"`

	// ID is the runtime ID that this policy is valid for.
	ID signature.PublicKey `json:"id"`

	// Enclaves is the per-key manager enclave ID access control policy.
	Enclaves map[sgx.EnclaveIdentity]*EnclavePolicySGX `json:"enclaves"`
}

// EnclavePolicySGX is the per-SGX key manager enclave ID access control policy.
type EnclavePolicySGX struct {
	// MayQuery is the map of runtime IDs to the vector of enclave IDs that
	// may query private key material.
	//
	// TODO: This could be made more sophisticated and seggregate based on
	// contract ID as well, but for now punt on the added complexity.
	MayQuery map[signature.MapKey][]sgx.EnclaveIdentity `json:"may_query"`

	// MayReplicate is the vector of enclave IDs that may retrieve the master
	// secret (Note: Each enclave ID may always implicitly replicate from other
	// instances of itself).
	MayReplicate []sgx.EnclaveIdentity `json:"may_replicate"`
}

// SignedPolicySGX is a signed SGX key manager access control policy.
type SignedPolicySGX struct {
	Policy PolicySGX `json:"policy"`

	Signatures []signature.Signature `json:"signatures"`
}
