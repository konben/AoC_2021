(* AoC Day 1: Sonar Sweep in sml-nj *)

(* val measurements = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]; *)
use "measurements.sml";

(* Counts how many measurements are greater than the previous. *)
fun countGreater (xs: int list) = countHelper 0 NONE xs
and countHelper n _ nil = n
|   countHelper n NONE (x::xs) = countHelper n (SOME x) xs
|   countHelper n (SOME last) (x::xs) =
        if x > last then countHelper (n + 1) (SOME x) xs
        else countHelper n (SOME x) xs;

val result = countGreater measurements;
print ((Int.toString result) ^ " results where greater than the previous!\n");
