# Copyright 2016 Markus W Mahlberg <markus.mahlberg@me.com>
#
# This file is part of cayleybench.
#
# cayleybench is free software: you can redistribute it and/or modify
# it under the terms of the GNU Lesser General Public License as published
# by the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# cayleybench is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Lesser General Public License for more details.
#
# You should have received a copy of the GNU Lesser General Public License
# along with cayleybench.  If not, see <http://www.gnu.org/licenses/>.

FROM alpine
MAINTAINER Markus W Mahlberg <markus.mahlberg@me.com>
ADD cayleybench.test /bin/cayleybench.test
ENTRYPOINT ["/bin/cayleybench.test", "-test.run=XXX" ,"-test.bench=.", "-test.benchmem","-sleep=5"]
CMD ["-test.cpu","1,2,4,6,8"]
