function request()
    wrk.method = "GET"

    path = "/" .. tostring(math.random(1000, 10000))

    return wrk.format(nil, path)
end
