Adding a graceful shutdown feature for server is going to address two problems :

- in-flight requests will be handled gracefully.
- shutdown can be looged.
- any work being handled by our handlers can be handled gracefully and can be completed and be shotdown gracefully.
