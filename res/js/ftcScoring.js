function isValidDate(date) {
    var temp = date.split('/');
    if(temp[0].length == 2 && Number(temp[0])<13 && Number(temp[0])>0){
        if(temp[1].length == 2 && Number(temp[1])<32 && Number(temp[1])>0){
            if(temp[2].length == 4&& Number(temp[2]) > 0){
                return true;
            }
            else  return false;
        }
        else  return false;
    }
    else  return false;
}