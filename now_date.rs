use time;
extern crate time;

fn main() {
    let now = time::now();
    let f_now = time::strftime("%Y-%m-%d", &now).unwrap();
    println!("now: {:?}", f_now)
}