These dockerfiles are planned for splitting the caddy and wfn images. This
change will need to wait until I figure out how to do networking in AWS's
Elastic Container Service (so caddy can proxy over port 8081 to wfn).
Until this I will be shimming the two processes into the same image with
the wfn.sh script as the entrypoint.
