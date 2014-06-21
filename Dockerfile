# Copyright 2014 Alexey Kryuchkov
#
# This file is part of tarserver.
#
# tarserver is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# tarserver is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with tarserver.  If not, see <http://www.gnu.org/licenses/>.
#
FROM ubuntu
ADD . /tarserver
WORKDIR tarserver
EXPOSE 80
ENTRYPOINT ["bin/tarserver"]
