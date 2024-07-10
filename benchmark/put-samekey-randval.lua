function request()
    wrk.method = "PUT"
    wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"
    wrk.body = "value=" .. tostring(math.random(1000, 10000))

    return wrk.format(nil)
end
