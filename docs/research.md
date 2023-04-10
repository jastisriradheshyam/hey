# Research

## Clap

- Benchmarks : https://github.com/rosetta-rs/argparse-rosetta-rs
- FAQ: https://docs.rs/clap/latest/clap/_faq/index.html
- Derive API tutorial: https://docs.rs/clap/latest/clap/_derive/_tutorial/index.html
- Builder API tutorial: https://docs.rs/clap/latest/clap/_tutorial/index.html
- Clap features: https://docs.rs/clap/latest/clap/_features/index.html

- Should I use Derive API or Builder API ?
  - Go for Derive API
    - Development is easy and faster with Derive API
    - If require more granular control over few features there is an option to do interop with Builder API
      - https://docs.rs/clap/latest/clap/_derive/index.html#mixing-builder-and-derive-apis

## System Info

- OS
  - (OS in std::env::consts - Rust)[https://doc.rust-lang.org/std/env/consts/constant.OS.html]
- ARCH
  - (ARCH in std::env::consts - Rust)[https://doc.rust-lang.org/std/env/consts/constant.ARCH.html]