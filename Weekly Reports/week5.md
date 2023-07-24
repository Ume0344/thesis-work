#### How to install [t4p4s](https://github.com/P4ELTE/t4p4s)

1. Install build-essential package `sudo apt-get install build-essential`

2. Make sure re2 library is installed. If not, do the following steps;
    - [re2-git-repo](https://github.com/google/re2)
    - git clone https://code.googlesource.com/re2
    - cd re2/
    - make
    - make test
    - make install
    - make testinstall

3. Run `. ./bootstrap-t4p4s.sh` on home directory of server.

4. It will successfully install the t4p4s with all dependencies. 

5. Run `cd t4p4s && ./t4p4s.sh %l2fwd` to confirm the installation. If the output is `T₄P₄S switch running ./examples/p4_14/l2fwd.p4_14  exited normally`, the t4p4s is installed successfully.
