FROM scratch
ADD milkMoneyBackend.tar.gz /etc
CMD ["/etc/milkMoneyBackend"]
