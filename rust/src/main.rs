use clap::Parser;
mod system_info;

// CLI help statement
/// Hey CLI Help
#[derive(Parser, Debug)]
#[command(author = "Jasti Sri Radhe Shyam <samabhsatejsrs@outlook.com>")]
#[command(about = "Cli tool for executing custom operations")]
#[command(version, long_about = None)]
struct Cli {
    /// System information
    #[arg(long, action)]
    system_info: bool,
}

fn main() {
    let args = Cli::parse();

    if args.system_info == true {
      system_info::get_system_info();
    }
}
