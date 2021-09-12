uAdmin Tutorial Part 3 - Image Cropping (Current Progress)
==========================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part3.html

Structure:

* `models`_
    * `category.go`_
    * `todo.go`_
* `main.go`_
* `uadmin.db`_
    * `Categories`_
    * `Todos`_

models
------

**category.go**
^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // Category Model !
    type Category struct {
        uadmin.Model
        Name string `uadmin:"required"`
        Icon string `uadmin:"image"`
    }

**todo.go**
^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "time"

        "github.com/uadmin/uadmin"
    )

    // Todo Model !
    type Todo struct {
        uadmin.Model
        Name        string
        Description string `uadmin:"html"`
        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

main.go
-------
`Back to Top`_

.. code-block:: go

    package main

    import (
        // Specify the username that you used inside github.com folder
        "github.com/username/todo/models"

        "github.com/uadmin/uadmin"
    )

    func main() {
        // Register Models
        uadmin.Register(
            models.Todo{},
            models.Category{},
        )

        uadmin.StartServer()
    }

uadmin.db
---------

**Categories**
^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/categorymodelupdate.png

**Todos**
^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part3.html#uadmin-tutorial-part-3-image-cropping-current-progress

.. image:: ../../tutorial/assets/todomodeloutput.png
