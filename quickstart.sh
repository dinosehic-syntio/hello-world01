#!/bin/sh
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

echo "Hello, world! The time is $(date)."
brew install bash-completion@2
echo 'export BASH_COMPLETION_COMPAT_DIR="/usr/local/etc/bash_completion.d"' >>~/.bashrc
echo '[[ -r "/usr/local/etc/profile.d/bash_completion.sh" ]] && . "/usr/local/etc/profile.d/bash_completion.sh"' >>~/.bashrc
exec bash # reload and replace (if it was updated) shell
type _init_completion && echo "completion is OK" # verify that bash-completion v2 is correctly installed