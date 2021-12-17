use argon2::{self};
use std::time::{Instant};
use threadpool::ThreadPool;


fn main(){
    let now = Instant::now();
    spawn_crackers();
    println!("Done in {}", now.elapsed().as_secs());
}


fn spawn_crackers() {
    let cpus = num_cpus::get();
    let pool = ThreadPool::new(cpus * 2);

    for i in 0..100000 {
        pool.execute(move || {
            cracker(i);
        });
    }
    pool.join();
}

fn cracker(i: i32){
    let mut password = "password".to_owned();
    password.push_str(&i.to_string());
    let hash = "$argon2id$v=19$m=16,t=2,p=1$MTExMTExMTE$SzwiO6Uix4CqutzHAncBwQ";
    argon2::verify_encoded(&hash, password.as_bytes()).unwrap();
    print!(".\n");
}
