{
  description = "TUI torrent search client using Knaben Database API";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";

  outputs = {
    self,
    nixpkgs,
  }: let
    systems = ["x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin"];
    forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f nixpkgs.legacyPackages.${system});
  in {
    packages = forAllSystems (pkgs: {
      default = pkgs.buildGoModule {
        pname = "torpar";
        version = "0.1.0";
        src = ./.;
        # Run `nix build` once; it will fail with the correct hash to paste here.
        vendorHash = pkgs.lib.fakeHash;
      };
    });

    devShells = forAllSystems (pkgs: {
      default = pkgs.mkShell {
        packages = [
          pkgs.go
          pkgs.gopls
          pkgs.alejandra # nix formatter
          # gofmt ships with go, no separate package needed
        ];
        shellHook = ''
          echo "Formatters: alejandra (nix), gofmt (go)"
        '';
      };
    });

    formatter = forAllSystems (pkgs: pkgs.alejandra);
  };
}
