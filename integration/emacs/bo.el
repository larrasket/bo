;; Copyright (C) 2023 Salih Muhammed
;;
;; Author: Salih Muhammed <salhghd7@gmail.com>
;; Maintainer: Salih Muhammed <salhghd7@gmail.com>
;; Created: April 07, 2023
;; Modified: April 07, 2023
;; Version: 0.0.1
;; Homepage: https://github.com/salehmu/bo
;; Package-Requires: ((emacs "24.3"))
;;
;; This file is not part of GNU Emacs.
;;
;;; Code:





(defun bo-trash-selected-items ()
  "Move selected items to trash using `trash-put`."
  (interactive)
  (let ((selected-items (dired-get-marked-files)))
    (dolist (item selected-items)
      (call-process "trash-put" nil nil nil item))
    (revert-buffer)))


(defun bo-trash-selected-items ()
  "Move selected items to trash using `trash-put`."
  (interactive)
  (let ((selected-items (dired-get-marked-files)))
    (dolist (item selected-items)
      (let ((command-string (format "trash-put %s" (shell-quote-argument item))))
        (shell-command command-string)))
    (revert-buffer)))


(defvar bo-dired-saved-dir nil
  "Variable to save the current directory.")

(defun bo-toggle ()
  "Toggle between the current directory and /tmp/boBrowser."
  (interactive)
  (if bo-dired-saved-dir
      (progn
        (dired bo-dired-saved-dir)
        (setq bo-dired-saved-dir nil))
    (progn
      (shell-command "bo")
      (setq bo-dired-saved-dir default-directory)
      (dired "/tmp/boBrowser"))))




(provide 'bo)
;;; bo.el ends here
