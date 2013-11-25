/*                                                                              
 * Copyright (C) 2013 Deepin, Inc.                                                 
 *               2013 Leslie Zhai <zhaixiang@linuxdeepin.com>                   
 *                                                                              
 * This program is free software: you can redistribute it and/or modify         
 * it under the terms of the GNU General Public License as published by         
 * the Free Software Foundation, either version 3 of the License, or            
 * any later version.                                                           
 *                                                                              
 * This program is distributed in the hope that it will be useful,              
 * but WITHOUT ANY WARRANTY; without even the implied warranty of               
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the                
 * GNU General Public License for more details.                                 
 *                                                                              
 * You should have received a copy of the GNU General Public License            
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.        
 */

package main

import (
    "fmt"
    "os"
    "github.com/xiangzhai/goftp"
)

func main() {
    // new ftp
    ftp := new(ftp.FTP)
    // set debug, default false
    ftp.Debug = true
    // connect
    ftp.Connect("localhost", 21)
    // login
    ftp.Login("anonymous", "")
    // login failure
    if ftp.Code == 530 {
        fmt.Println("error: login failure")
        os.Exit(-1)
    }
    // pwd
    ftp.Pwd()
    fmt.Println("code:", ftp.Code, ", message:", ftp.Message)
    ftp.Request("LS ")
    fmt.Println("code:", ftp.Code, ", message:", ftp.Message)
    // quit
    ftp.Quit()
}