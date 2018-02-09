


    


        

            
# Front End Development[¶](#front-end-development)

            

                
## Background[¶](#background)

                

                    
Note

                    
Consider this the canonical resource for contributing Javascript and CSS. We
                        are currently in the process of modernizing our front end development
                        procedures. You will see a lot of different styles around the code base for
                        front end JavaScript and CSS.

                

                
Our modern front end development stack includes the following tools:

                
- [Gulp](http://gulpjs.com)
- [Bower](http://bower.io)
- [Browserify](http://browserify.org)
- [Debowerify](https://github.com/eugeneware/debowerify)
- Andsoon,[LESS](http://lesscss.org)

                
We use the following libraries:

                
- [Knockout](http://knockoutjs.com)
- [jQuery](http://jquery.com)
- SeveraljQueryplugins

                
Previously, JavaScript development has been done in monolithic files or inside
                    templates. jQuery was added as a global object via an include in the base
                    template to an external source. There are no standards currently to JavaScript
                    libraries, this aims to solve that.

                
The requirements for modernizing our front end code are:

                
- Codeshouldbemodularandtestable.One-offchunksofJavaScriptintemplates
orinlargemonolithicfilesarenoteasilytestable.Wecurrentlyhaveno
JavaScripttests.
- Reducecodeduplication.
- EasyJavaScriptdependencymanagement.

                
Modularizing code with [Browserify](http://browserify.org) is a good first step. In this development
                    workflow, major dependencies commonly used across JavaScript includes are
                    installed with [Bower](http://bower.io) for testing, and vendorized as standalone libraries via
                    [Gulp](http://gulpjs.com) and [Browserify](http://browserify.org). This way, we can easily test our JavaScript libraries
                    against jQuery/etc, and have the flexibility of modularizing our code. See
                    [JavaScript Bundles](#javascript-bundles) for more information on what and how we are bundling.

                
To ease deployment and contributions, bundled JavaScript is checked into the
                    repository for now. This ensures new contributors don’t need an additional front
                    end stack just for making changes to our Python code base. In the future, this
                    may change, so that assets are compiled before deployment, however as our front
                    end assets are in a state of flux, it’s easier to keep absolute sources checked
                    in.

            

            

                
## Getting Started[¶](#getting-started)

                
You will need a working version of Node and NPM to get started. We won’t cover
                    that here, as it varies from platform to platform.

                
To install these tools and dependencies:

                


<pre>npm install
</pre>


                

                
This will install locally to the project, not globally. You can install globally
                    if you wish, otherwise make sure <code class="docutils literal">node_modules/.bin</code> is in your PATH.

                
Next, install front end dependencies:

                


<pre>bower install
</pre>


                

                
The sources for our bundles are found in the per-application path
                    <code class="docutils literal">static-src</code>, which has the same directory structure as <code class="docutils literal">static</code>. Files in
                    <code class="docutils literal">static-src</code> are compiled to <code class="docutils literal">static</code> for static file collection in Django.
                    Don’t edit files in <code class="docutils literal">static</code> directly, unless you are sure there isn’t a
                    source file that will compile over your changes.

                
To test changes while developing, which will watch source files for changes and
                    compile as necessary, you can run [Gulp](http://gulpjs.com) with our development target:

                


<pre>gulp dev
</pre>


                

                
Once you are satisfied with your changes, finalize the bundles (this will
                    minify library sources):

                


<pre>gulp build
</pre>


                

                
If you updated any of our vendor libraries, compile those:

                


<pre>gulp vendor
</pre>


                

                
Make sure to check in both files under <code class="docutils literal">static</code> and <code class="docutils literal">static-src</code>.

            

            

                
## Making Changes[¶](#making-changes)

                
If you are creating a new library, or a new library entry point, make sure to
                    define the application source file in <code class="docutils literal">gulpfile.js</code>, this is not handled
                    automatically right now.

                
If you are bringing in a new vendor library, make sure to define the bundles you
                    are going to create in <code class="docutils literal">gulpfile.js</code> as well.

                
Tests should be included per-application, in a path called <code class="docutils literal">tests</code>, under the
                    <code class="docutils literal">static-src/js</code> path you are working in. Currently, we still need a test
                    runner that accumulates these files.

            

            

                
## Deployment[¶](#deployment)

                
If merging several branches with JavaScript changes, it’s important to do a
                    final post-merge bundle. Follow the steps above to rebundle the libraries, and
                    check in any changed libraries.

            

            

                
## JavaScript Bundles[¶](#javascript-bundles)

                
There are several components to our bundling scheme:

                
<blockquote>
                    
<dl class="docutils">
                        <dt>Vendor library</dt>
                        <dd>
We repackage these using [Browserify](http://browserify.org), [Bower](http://bower.io), and [Debowerify](https://github.com/eugeneware/debowerify) to
                            make these libraries available by a <code class="docutils literal">require</code> statement.  Vendor
                            libraries are packaged separately from our JavaScript libraries, because
                            we use the vendor libraries in multiple locations. Libraries bundled
                            this way with [Browserify](http://browserify.org) are available to our libraries via
                            <code class="docutils literal">require</code> and will back down to finding the object on the global
                            <code class="docutils literal">window</code> scope.

                            
Vendor libraries should only include libraries we are commonly reusing.
                                This currently includes <code class="xref py py-obj docutils literal">jQuery</code> and <code class="xref py py-obj docutils literal">Knockout</code>. These modules will be
                                excluded from libraries by special includes in our <code class="docutils literal">gulpfile.js</code>.

                        </dd>
                        <dt>Minor third party libraries</dt>
                        <dd>These libraries are maybe used in one or two locations. They are
                            installed via [Bower](http://bower.io) and included in the output library file. Because
                            we aren’t reusing them commonly, they don’t require a separate bundle or
                            separate include. Examples here would include jQuery plugins used on one
                            off forms, such as jQuery Payments.</dd>
                        <dt>Our libraries</dt>
                        <dd>
These libraries are bundled up excluding vendor libraries ignored by
                            rules in our <code class="docutils literal">gulpfile.js</code>. These files should be organized by
                            function and can be split up into multiple files per application.

                            
Entry points to libraries must be defined in <code class="docutils literal">gulpfile.js</code> for now. We
                                don’t have a defined directory structure that would make it easy to
                                imply the entry point to an application library.

                        </dd>
                    </dl>
                    
</blockquote>

            

        



    

    


    




