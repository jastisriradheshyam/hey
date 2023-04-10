use clap::Parser;
mod system_info;

// CLI help statement
/// Hey CLI Help
#[derive(Parser, Debug)]
#[command(author = "Jasti Sri Radhe Shyam <samabhsatejsrs@outlook.com>")]
#[command(about = "Cli tool for executing custom operations")]
#[command(version, long_about = None)]
struct Cli {
    /// Name of the person to greet
    #[arg(short, long)]
    name: String,

    /// Number of times to greet
    #[arg(short, long, default_value_t = 1)]
    count: u8,

    // System information
    #[arg(long, action)]
    system_info: bool,
}

fn main() {
    let args = Cli::parse();

    for _ in 0..args.count {
        println!("Hello {}!", args.name)
    }

    if args.system_info == true {
      system_info::get_system_info();
    }
}
