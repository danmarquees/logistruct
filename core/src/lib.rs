#[no_mangle]
pub extern "C" fn
calculate_safety_stock(
    max_sales: i32,
    max_lead_time: i32,
    avg_sales: i32,
    avg_lead_time: i32,
) -> i32 {
    let safety_stock = (max_sales * max_lead_time) - (avg_sales * avg_lead_time);
    safety_stock
}
