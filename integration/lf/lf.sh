# Trash cli bindings
cmd trash ${{
  files=$(printf "$fx" | tr '\n' ';')
  while [ "$files" ]; do
    # extract the substring from start of string up to delimiter.
    # this is the first "element" of the string.
    file=${files%%;*}

    trash-put "$(basename "$file")"
    # if there's only one element left, set `files` to an empty string.
    # this causes us to exit this `while` loop.
    # else, we delete the first "element" of the string from files, and move onto the next.
    if [ "$files" = "$file" ]; then
      files=''
    else
      files="${files#*;}"
    fi
  done
}}

cmd clear_trash %trash-empty

cmd bo_start ${{
bo
boloc="/tmp/boBrowser"
cmd="cd"
lf -remote "send $id $cmd \"$boloc\""

}}



cmd bo_restore ${{
  files=$(printf "$fx" | tr '\n' ' ')
  /home/ghd/go/bin/bo r $files
}}

cmd bo_getback ${{
  file_path=$(cat /tmp/boBrowser/.bin_organizer12)
  cmd="cd"
  lf -remote "send $id $cmd \"$file_path\""
}}


# Trash Mappings
map dd trash
map ds bo_start
map dr bo_restore
map g` bo_getback
