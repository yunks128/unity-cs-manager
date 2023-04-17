{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = [
    pkgs.go_1_18
    pkgs.gotools
    pkgs.gopls
    pkgs.go-outline
    pkgs.gocode
    pkgs.gopkgs
    pkgs.gocode-gomod
    pkgs.godef
    pkgs.golint
    pkgs.go-tools
    pkgs.golangci-lint
    pkgs.gofumpt
    pkgs.pre-commit
    pkgs.gosec
    pkgs.nodejs
    pkgs.nodePackages.npm
    pkgs.ruby];
}
