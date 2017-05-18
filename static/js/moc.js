class MOC {
    formatTime(t) {
        var time = moment(t)
        return time.format("YYYY-MM-DD HH:mm:ss")
    }
    formatTimeRange(t){
        var time = moment(t)
        return time.format("HH:mm")
    }
    toggleLoading(flag){
        if (flag){
            $("#loading").css("display","block")
        }else{
            $("#loading").css("display","none")
        }
    }
}

var moc = new(MOC);

