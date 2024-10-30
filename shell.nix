{ pkgs ? import <nixos> {} }:
pkgs.mkShell {
  nativeBuildInputs = with pkgs; [
    gcc
    go
  ];
}
