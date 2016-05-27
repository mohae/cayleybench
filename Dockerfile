FROM alpine
ADD cayleybench.test /bin
ENTRYPOINT ["/bin/cayleybench.test", "-test.run=XXX" ,"-test.bench=.", "-test.benchmem","-sleep=5"]
CMD ["-test.cpu","1,2,4,6,8"]
