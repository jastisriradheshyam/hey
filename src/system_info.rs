use std::env;

pub fn get_system_info() {
    println!("OS: {}", env::consts::OS);
    println!("ARCH: {}", env::consts::ARCH);
}
