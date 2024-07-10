function request()
    wrk.method = "PUT"
    wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"
    wrk.body = "value=" .. tostring(math.random(1000, 1000000))

    return wrk.format()
end

-- function request()
--     counter = counter + 1

--     return wrk.request()
-- end
