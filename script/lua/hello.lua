local orm = require('orm')

function success()
    outData['code'] = 1
    outData['msg'] = 'SUCCESS'
end

function fail(code, err)
    outData['code'] = code
    outData['msg'] = err
end

function get_aql_list()
    local sql = [[
        select * from qc_aql
    ]]
    local t = orm.query(sql)
    outData['data'] = t
    return t == nil and fail(-1, orm.err()) or success()
end

function get_aql_level_list()
    local sql = [[
        select * from qc_aql_level where f_qc_aql_level_id < ?
    ]]
    local tParam = {}
    tParam[1] = 3
    local t = orm.query(sql, tParam)
    outData['data'] = t
    return t == nil and fail(-1, orm.err()) or success()
end

function process()
    local cmd = inData['cmd']
    if _G[cmd] ~= nil then
        _G[cmd]()
    else
        fail(-1, '接口不存在')
    end
end

process()