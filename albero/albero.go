package albero

import (
	"bytes"
	"encoding/hex"
)

// Defines the Tree structure
type Tree struct {
	Root   *Node
	Leaves []*Node
}

// Defines the Node structure
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Hash   []byte
}

// Helper function for hex encoding
func (n *Node) Hex() string {
	// Returns the encoded string in the hex format
	return hex.EncodeToString(n.Hash)
}

// Creates a Merkle tree from input hashes
func NewTreeFromHashes(hashes [][]byte) *Tree {

	// Gets the length of the hashes
	l := len(hashes)

	// Performs integrity checks on input data
	if l == 0 {
		panic(ErrInvalidInput)
	}

	// Normalizes potential even leafs
	if l%2 != 0 {
		if l > 1 {
			hashes = append(hashes, hashes[l-1])
		}
	}

	// Generates tree
	t := &Tree{
		Leaves: make([]*Node, 0, len(hashes)),
	}

	// Creates leaf nodes
	for _, h := range hashes {
		t.Leaves = append(t.Leaves, &Node{Hash: h})
	}

	// Builds and sets the root
	t.Root = t.buildRoot()

	// Returns the tree
	return t
}

// Verifies the integrity of a given input value
func VerifyMerkleProof(rootHash, value []byte, proofs [][]byte, idxs []int) bool {
	prevHash := value
	// Loops through the proofs
	for i := 0; i < len(proofs); i++ {
		if idxs[i] == 0 {
			prevHash = PerformHash(append(proofs[i], prevHash...))
		} else {
			prevHash = PerformHash(append(prevHash, proofs[i]...))
		}
	}

	return bytes.Equal(rootHash, prevHash)
}

// Returns the Merkle path proof to verify the integrity of the given input hash
func (t *Tree) GenerateMerkleProof(hash []byte) ([][]byte, []int, error) {
	// Defines implied variables
	var (
		path [][]byte
		idxs []int
	)

	// Finds the leaf node for the specific hash
	for _, currentNode := range t.Leaves {
		if bytes.Equal(currentNode.Hash, hash) {
			// Scales the tree using the relationship of the nodes to their parent nodes
			parent := currentNode.Parent
			for parent != nil {
				// In order to calculate the parent hash for a proof in a Merkle tree, we need to consider the position of the current node in relation to its parent
				// If the current node is the left child, we require the right child's hash to calculate the parent hash for the proof
				// Conversely, if the current node is the right child, we need the left child's hash
				// - If the current node is the left child: ParentHash = (CurrentNode.Hash, RightChild.Hash)
				// - If the current node is the right child: ParentHash = (LeftChild.Hash, CurrentNode.Hash)
				// Therefore, when constructing the proof, we add the corresponding hash to the proof path
				// Additionally, in the "idxs" list, we save the position of the hash: 0 for the left child and 1 for the right child
				// This allows us to determine, during the verification of the proof, whether the given hash is the left or right child based on its position in the proof path
				if bytes.Equal(currentNode.Hash, parent.Left.Hash) {
					path = append(path, parent.Right.Hash)
					idxs = append(idxs, 1)
				} else {
					path = append(path, parent.Left.Hash)
					idxs = append(idxs, 0)
				}
				currentNode = parent
				parent = currentNode.Parent
			}
			return path, idxs, nil
		}
	}
	return path, idxs, ErrInvalidProof
}

// Rebuilds the tree to verify its integrity
func (t *Tree) VerifyTreeIntegrity() bool {
	// Handles the basic case
	if len(t.Leaves) == 0 || t.Root == nil {
		return false
	}

	r := t.buildRoot()
	return bytes.Equal(t.Root.Hash, r.Hash)
}

// Effectively builds the root of the Merkle tree
func (t *Tree) buildRoot() *Node {
	nodes := t.Leaves
	// We are iterating until we reach a single node, which will be our root

	var parents []*Node

	// Pairing nodes to build a parent from the pair
	for i := 0; i < len(nodes); i += 2 {
		n := &Node{
			Left:  nodes[i],
			Right: nodes[i+1],

			// Compute the hash of the new node, which will be the combination of its children's hashes
			Hash: PerformHash(append(nodes[i].Hash, nodes[i+1].Hash...)),
		}

		parents = append(parents, n)
		nodes[i].Parent, nodes[i+1].Parent = n, n
	}
	// Once all possible pairs are processed, the parents become the children, and we start all over again
	nodes = parents

	// Returns the root
	return nodes[0]
}
