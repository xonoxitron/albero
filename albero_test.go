package main

import (
	"bufio"
	"consensys/home-assignment/albero"
	"os"
	"testing"
)

// Hosts mocked transactions
var mockTxs []string

// Loads transactions in the global variable
func setup() {
	numRecords := os.Args[len(os.Args)-1]
	filePath := "test/transactions_" + numRecords + ".txt"
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mockTxs = append(mockTxs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// Cleans up resources more needed
func teardown() {
	mockTxs = nil
}

// Pre-test + Tests orchestration
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// Creates new tree from mocked input transactions
func createTree() *albero.Tree {
	// Prepares the input data
	inputData := make([][]byte, len(mockTxs))

	// Performs input data hashing
	for i, str := range mockTxs {
		inputData[i] = albero.PerformHash([]byte(str))
	}

	// Generates the Merkle tree from the input data
	tree := albero.NewTreeFromHashes(inputData)
	return tree
}

// Tests the Merkle Tree creation
func TestMerkleTreeCreation(t *testing.T) {

	inputData := [][]byte{
		albero.PerformHash([]byte("From: 0x108a62c6cd1c2a84c17736e52c7d92960b26af55, To: 0x081c215b9665db6cf1b46414b95adc193dc2b8f6, ETH: 0.668486374749")),
		albero.PerformHash([]byte("From: 0xb09387d76772f67fde196bedef44609b2e98d03e, To: 0x3798ecdb690bb1032fe1272ab9ca7a9237b87217, ETH: 0.970741064563")),
		albero.PerformHash([]byte("From: 0x09efd65cb2c98b94c581f625fc35ad1929ba6298, To: 0xfd7f40362edb56f1505e2886d40911d72e556c50, ETH: 0.850225331564")),
	}

	tree := albero.NewTreeFromHashes(inputData)
	treeRootHex := tree.Root.Hex()

	// Final check
	if treeRootHex != "758647d30d5e520f2ff5a3eb09507c96e6480a2491c01dda0e20ab7d9d9b2a80" {
		t.Fatalf("fail")
	}
}

// Tests the Merkle Proof integrity
func TestMerkleProofIntegrity(t *testing.T) {
	// Creates new tree
	tree := createTree()

	// Gets the proof of the first transaction and verify it.
	tx1Hash := albero.PerformHash([]byte(mockTxs[0]))
	proof, idxs, err := tree.GenerateMerkleProof(tx1Hash)
	if err != nil {
		t.Fatalf("fail")
	}
	// Gets the Merkle Proof integrity response
	proofIntegrity := albero.VerifyMerkleProof(tree.Root.Hash, tx1Hash, proof, idxs)

	// Final check
	if !proofIntegrity {
		t.Fatalf("fail")
	}
}

// Tests the Merkle Tree integrity
func TestMerkleTreeIntegrity(t *testing.T) {

	// Creates new tree
	tree := createTree()

	// Verifies again the integrity of the tree, before the modification
	treeIntegrity := tree.VerifyTreeIntegrity()

	// Integrity is TRUE
	if !treeIntegrity {
		t.Fail()
	}

	// Simulates an integrity break by modifiyng a transaction
	tree.Leaves[0].Hash = albero.PerformHash([]byte("From: 0x006546305b8eab155bf087b176fad3cf45ecaf1a, To: 0x30f7225274fba7d6b431a11eabf02556407ee439, ETH: 0.200815824264"))

	// Verifies again the integrity of the tree, after the modification
	treeIntegrity = tree.VerifyTreeIntegrity()

	// Integrity is FALSE
	if treeIntegrity {
		t.Fail()
	}
}
