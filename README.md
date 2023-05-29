# 🌳 albero: Merkle Tree Implementation in Go 🐿️

Albero is a Go library that provides an implementation of a Merkle tree along with functionalities to support Merkle proof integrity verification and Merkle tree integrity verification. This library is designed to be simple, efficient, and easy to use.

## Table of Contents

- [🌳 albero: Merkle Tree Implementation in Go 🐿️](#-albero-merkle-tree-implementation-in-go-️)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Creating a Merkle Tree](#creating-a-merkle-tree)
    - [Verifying Merkle Proof](#verifying-merkle-proof)
    - [Verifying Merkle Tree Integrity](#verifying-merkle-tree-integrity)
  - [Benchmark](#benchmark)
    - [Testing hardware](#testing-hardware)
      - [1'000 Samples RUN](#1000-samples-run)
      - [10'000 Samples RUN](#10000-samples-run)
      - [100'000 Samples RUN](#100000-samples-run)
      - [1'000'000 Samples RUN](#1000000-samples-run)
  - [Contributing](#contributing)
  - [License](#license)

## Introduction

A Merkle tree, also known as a hash tree, is a data structure that allows efficient verification of the integrity and consistency of a large amount of data. It achieves this by organizing the data into a binary tree structure, where each leaf node represents a data block, and each non-leaf node represents the hash of its child nodes.

The "albero" library provides an easy-to-use interface to create and work with Merkle trees. It supports generating Merkle proofs for individual data blocks, as well as verifying the integrity of the entire tree.

## Installation

To use the "albero" library in your Go project, you need to have Go installed and set up on your system. Once you have Go installed, you can install the library by running the following command:

```bash
go get github.com/xonoxitron/albero
```

Make sure to replace `xonoxitron` with your actual GitHub username or the appropriate repository URL.

## Usage

### Creating a Merkle Tree

To create a Merkle tree using the "albero" library, you first need to import the package:

```go
import "github.com/xonoxitron/albero"
```

Then, you can use the `NewTreeFromHashes` function to create a Merkle tree from a list of data hashes:

```go
hashes := [][]byte{hash1, hash2, hash3, ...}
tree := albero.NewTreeFromHashes(hashes)
```

The `hashes` parameter should be a slice of byte slices, where each byte slice represents the hash of a data block. The function returns a `Tree` object representing the created Merkle tree.

### Verifying Merkle Proof

To verify the integrity of a specific data block using a Merkle proof, you can use the `VerifyMerkleProof` function:

```go
rootHash := []byte("root hash of the Merkle tree")
value := []byte("data block hash")
proofs := [][]byte{proof1, proof2, proof3, ...}
indexes := []int{index1, index2, index3, ...}

isValid := albero.VerifyMerkleProof(rootHash, value, proofs, indexes)
```

The `rootHash` parameter should be the root hash of the Merkle tree, `value` should be the hash of the data block you want to verify, `proofs` should be a slice of byte slices representing the Merkle proof path, and `indexes` should be a slice of integers indicating the position of each proof hash (0 for the left child, 1 for the right child). The function returns a boolean value indicating whether the verification is successful or not.

### Verifying Merkle Tree Integrity

To verify the integrity of the entire Merkle tree, you can use the `VerifyTree

Integrity` method of the `Tree` object:

```go
isValid := tree.VerifyTreeIntegrity()
```

The function returns a boolean value indicating whether the Merkle tree is intact or not.

## Benchmark

### Testing hardware

- OS: Venture 13.3.1
- CPU: 2,3 GHz 8-Core Intel Core i9
- RAM: 16 GB

#### 1'000 Samples RUN

```console
(base) matteo@Matteos-MacBook-Pro consensys-home-assignment % sh run.sh 1000 
Generating 1000 random transaction records...
Generated 1000 random transaction records and wrote them to transactions_1000.txt
=== RUN   TestMerkleTreeCreation
--- PASS: TestMerkleTreeCreation (0.00s)
=== RUN   TestMerkleProofIntegrity
--- PASS: TestMerkleProofIntegrity (0.00s)
=== RUN   TestMerkleTreeIntegrity
--- PASS: TestMerkleTreeIntegrity (0.00s)
PASS
ok      consensys/home-assignment       0.207s
```

#### 10'000 Samples RUN

```console
(base) matteo@Matteos-MacBook-Pro consensys-home-assignment % sh run.sh 10000 
Generating 10000 random transaction records...
Generated 10000 random transaction records and wrote them to transactions_10000.txt
=== RUN   TestMerkleTreeCreation
--- PASS: TestMerkleTreeCreation (0.00s)
=== RUN   TestMerkleProofIntegrity
--- PASS: TestMerkleProofIntegrity (0.01s)
=== RUN   TestMerkleTreeIntegrity
--- PASS: TestMerkleTreeIntegrity (0.01s)
PASS
ok      consensys/home-assignment       0.134s
```

#### 100'000 Samples RUN

```console
(base) matteo@Matteos-MacBook-Pro consensys-home-assignment % sh run.sh 100000
Generating 100000 random transaction records...
Generated 100000 random transaction records and wrote them to transactions_100000.txt
=== RUN   TestMerkleTreeCreation
--- PASS: TestMerkleTreeCreation (0.00s)
=== RUN   TestMerkleProofIntegrity
--- PASS: TestMerkleProofIntegrity (0.08s)
=== RUN   TestMerkleTreeIntegrity
--- PASS: TestMerkleTreeIntegrity (0.11s)
PASS
ok      consensys/home-assignment       0.324s
```

#### 1'000'000 Samples RUN

```console
(base) matteo@Matteos-MacBook-Pro consensys-home-assignment % sh run.sh 1000000
Generating 1000000 random transaction records...
Generated 1000000 random transaction records and wrote them to transactions_1000000.txt
=== RUN   TestMerkleTreeCreation
--- PASS: TestMerkleTreeCreation (0.00s)
=== RUN   TestMerkleProofIntegrity
--- PASS: TestMerkleProofIntegrity (0.72s)
=== RUN   TestMerkleTreeIntegrity
--- PASS: TestMerkleTreeIntegrity (1.15s)
PASS
ok      consensys/home-assignment       2.228s
```

The benchmark section provides information on the performance of the "albero" library on a specific hardware setup. It includes the hardware specifications, as well as the execution times for different sample sizes.

## Contributing

Contributions to the "albero" library are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/xonoxitron/albero).

## License

The "albero" library is licensed under the [MIT License](LICENSE). You can freely use, modify, and distribute the library, subject to the terms of the license.
