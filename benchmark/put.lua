wrk.method = "PUT"
wrk.body   = "value=" .. tostring(math.random(1000, 1000000))
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"
