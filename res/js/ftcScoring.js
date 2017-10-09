function isValidDate(date) {
    var temp = date.split('/');
    var d = new Date(temp[2] + '/' + temp[0] + '/' + temp[1]);
    return (d && (d.getMonth() + 1) == temp[0] && d.getDate() == Number(temp[1]) && d.getFullYear() == Number(temp[2]));
}