# Go File Encryption Tool

## Overview

The **Go File Encryption Tool** is a lightweight, command-line interface (CLI)-based application designed to securely encrypt and decrypt files. Developed in Go, it utilizes advanced cryptographic standards to ensure robust data confidentiality and integrity, making it ideal for protecting sensitive information.

---

## Features

- **AES-GCM Encryption**: Ensures strong data security with tamper detection.
- **Password-Based Protection**: Derives encryption keys using user-provided passwords.
- **File Integrity Validation**: Processes only valid and accessible files.
- **Cross-Platform CLI**: Seamless command-line experience for encryption and decryption tasks.

---

## Encryption Flow

1. **Password Input**:
   - The user provides a password via the CLI.
   - A confirmation step ensures password accuracy and avoids input errors.

2. **Key Derivation**:
   - The password is used with PBKDF2 (Password-Based Key Derivation Function 2) and SHA-1 to generate a 256-bit encryption key.
   - A unique 12-byte random nonce is created for each encryption operation to ensure ciphertext uniqueness.

3. **Encryption Process**:
   - The file’s plaintext data is read and encrypted using AES-GCM.
   - The resulting ciphertext, combined with the nonce, is written back to the file.

4. **Decryption Process**:
   - The nonce is extracted from the encrypted file.
   - Using the same password and nonce, the ciphertext is decrypted to retrieve the original plaintext.
   - The plaintext data is written back to the file, restoring its original state.

---

## Project Structure

### `main.go`

Handles the CLI interface, user input, and routing to encryption or decryption functionality.

- **Key Functions**:
  - `encryptHandle()`: Manages the file encryption process.
  - `decryptHandle()`: Manages the file decryption process.
  - `getPassword()`: Captures and validates the user’s password securely.

### `filecrypt.go`

Contains the core logic for encryption and decryption operations.

- **Key Functions**:
  - `Encrypt()`: Encrypts the file’s contents using AES-GCM and appends the nonce.
  - `Decrypt()`: Decrypts the file’s contents using the provided password and restores the original data.

---

## Usage Guide

### Commands

- **Encrypt a File**:

  ```bash
  go run . encrypt /path/to/file
  ```

  Prompts the user for a password and encrypts the specified file.

- **Decrypt a File**:

  ```bash
  go run . decrypt /path/to/file
  ```

  Prompts the user for the encryption password and decrypts the file.

- **Help**:

  ```bash
  go run . help
  ```

  Displays usage instructions and available commands.

---

## Encryption Methodology

### Algorithms

- **AES-GCM**:
  - Provides authenticated encryption to ensure confidentiality and integrity.
  - Tamper-evident ciphertext to detect unauthorized modifications.

- **PBKDF2**:
  - Generates a 256-bit encryption key from the user’s password.
  - Uses 4096 iterations to enhance resistance against brute-force attacks.

- **SHA-1 (in PBKDF2)**:
  - Employed as part of PBKDF2 for key derivation. Secure within this context despite deprecation for direct cryptographic use.

### Key Security Features

1. **Random Nonce**:
   - A unique 12-byte nonce ensures each encryption operation results in distinct ciphertext, even with the same input and password.

2. **Password Protection**:
   - The user’s password is integral to encryption and decryption, making it impossible to decrypt files without it.

3. **Data Integrity**:
   - AES-GCM ensures that tampered ciphertext cannot be decrypted successfully, safeguarding against unauthorized modifications.

---






## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to enhance functionality, improve security, or expand features.

---

## Disclaimer

This tool is intended for educational purposes and personal file security. Use it responsibly and at your own risk.

