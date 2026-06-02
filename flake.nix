{
  description = "ghq-wt - Manage remote repository clones with git worktree layout (binary: ghq)";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f nixpkgs.legacyPackages.${system});
    in
    {
      packages = forAllSystems (pkgs: rec {
        ghq-wt = pkgs.buildGoModule {
          pname = "ghq-wt";
          version = self.shortRev or "dev";
          src = self;
          vendorHash = "sha256-UddGBIjdUejoYT4L5ke8rDGd8Lp6/TM3QBEdi+DrbkQ=";

          ldflags = [
            "-s" "-w"
            "-X main.revision=${self.shortRev or "dev"}"
          ];

          # Tests require network access and git
          doCheck = false;

          # The binary is called "ghq"
          postInstall = ''
            mv $out/bin/ghq-wt $out/bin/ghq 2>/dev/null || true
          '';

          meta = {
            description = "Manage remote repository clones with git worktree layout";
            homepage = "https://github.com/gnur/ghq-wt";
            license = pkgs.lib.licenses.mit;
            mainProgram = "ghq";
          };
        };
        default = ghq-wt;
      });
    };
}
