function findMax(arr) {
    var max = arr[0];
    for (var _i = 0, arr_1 = arr; _i < arr_1.length; _i++) {
        var num = arr_1[_i];
        if (num > max) {
            max = num;
        }
    }
    return max;
}
