| GET / sec | PUT / sec | Concurrent  | Random Keys | Random Values | Key Validation | Get Mutex | Put Mutex |
| --------- | --------- | ----------- | ----------- | ------------- | -------------- | --------- | --------- |
| 103k      | 97.7k     | Yes         | Yes         | Yes           | Yes            | Yes       | Yes       |
| 175k      | 168k      | No          | Yes         | Yes           | Yes            | Yes       | Yes       |
| 105k      | 101k      | Yes         | Yes         | Yes           | No             | Yes       | Yes       |
| 181k      | 168k      | No          | Yes         | Yes           | No             | Yes       | Yes       |
| N/A       | N/A       | Yes         | Yes         | Yes           | Yes            | No        | Yes       |
| N/A       | N/A       | Yes         | Yes         | Yes           | Yes            | No        | No        |
| 180k      | 169k      | No          | Yes         | Yes           | Yes            | No        | Yes       |
| N/A       | -         | No          | Yes         | Yes           | Yes            | No        | No        |
