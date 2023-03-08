inData = JSON.parse(inData)
outData = {
    code: 1,
    msg: "SUCCESS",
    data: []
}

function get_aql_list() {
    var sql = "select * from qc_aql"
    var sJson = Orm.query(sql)
    var result = JSON.parse(sJson)
    if (!result["flag"]) {
        outData.code = -1
        outData.msg = result["err"]
    } else {
        outData.data = result["data"]
    }
}

function get_aql_level_list() {
    var sql = `select * from qc_aql_level where f_qc_aql_level_id < ?`
    var sJson = Orm.query(sql, 3)
    var result = JSON.parse(sJson)
    if (!result["flag"]) {
        outData.code = -1
        outData.msg = result["err"]
    } else {
        outData.data = result["data"]
    }
}

function process() {
    var cmd = inData["cmd"]
    if (typeof(eval(cmd)) == undefined) {
        outData["code"] = -1
        outData["msg"] = "接口不存在"
    } else {
        eval(cmd)()
    }
    outData = JSON.stringify(outData)
}

process()