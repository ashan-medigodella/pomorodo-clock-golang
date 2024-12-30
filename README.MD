# Pomodoro Clock Implementation in Go

This document explains the implementation of a Pomodoro Clock in Go, utilizing Windows notifications to enhance usability. The application manages work and break cycles using a timer, displays notifications, and handles file paths dynamically.

---

## **Overview**
The program implements a Pomodoro Clock with the following features:
- **Work and Break Cycles**: Configurable durations for work sessions and short/long breaks.
- **Windows Notifications**: Native notifications to inform the user of session changes.
- **Dynamic Path Handling**: Ensures the notification icon is loaded correctly from a relative path.

---

## **Program Flow**
1. **Initialize Constants**: Define durations for work sessions, short breaks, long breaks, and the number of cycles.
2. **Loop Through Cycles**: Alternate between work and break periods, showing notifications for each.
3. **Timer Functionality**: Use a countdown timer for each session.
4. **Notification Handling**: Display notifications with a title, message, and icon.

---

## **Code Explanation**

### **Constants**
```go
const workDuration = 25 * time.Minute
const shortBreak = 5 * time.Minute
const longBreak = 15 * time.Minute
const cycles = 4
```
These constants define the durations for work and break sessions, as well as the number of Pomodoro cycles.

---

### **Main Function**
```go
func main() {
    for i := 1; i <= cycles; i++ {
        fmt.Printf("Cycle %02d: Work for %02d minutes\n", i, workDuration/time.Minute)
        showNotification("Pomodoro Clock", fmt.Sprintf("Cycle %02d: Work for %02d minutes", i, workDuration/time.Minute))
        startTimer(workDuration)

        if i < cycles {
            shortBreakToString := fmt.Sprintf("Take a %02d minute break\n", shortBreak/time.Minute)
            showNotification("Pomodoro Clock", shortBreakToString)
            fmt.Println(shortBreakToString)
            startTimer(shortBreak)
        } else {
            longBreakToString := fmt.Sprintf("Take a %02d minute break\n", longBreak/time.Minute)
            showNotification("Pomodoro Clock", longBreakToString)
            fmt.Println(longBreakToString)
            startTimer(longBreak)
        }
    }
}
```
- **Work Cycles**: Alternates between work and break durations.
- **Notifications**: Displays a Windows notification for each session transition.
- **Timer Management**: Calls `startTimer` to manage countdowns.

---

### **Timer Function**
```go
func startTimer(duration time.Duration) {
    for duration > 0 {
        fmt.Printf("\rTime left: %02d:%02d", int(duration.Minutes()), int(duration.Seconds())%60)
        time.Sleep(1 * time.Second)
        duration -= time.Second
    }
    fmt.Println("\nTime's up!")
}
```
- **Countdown Loop**: Decrements the timer every second and displays the remaining time in minutes and seconds.
- **Completion Message**: Prints "Time's up!" when the timer ends.

---

### **Notification Function**
```go
func showNotification(title, message string) {
    baseDir, err := os.Getwd()
    if err != nil {
        fmt.Printf("Error getting current directory: %v", err)
        return
    }

    imgPath := filepath.Join(baseDir, "assets", "clock.png")

    notification := toast.Notification{
        AppID:   "Pomodoro Clock",
        Title:   title,
        Message: message,
        Icon:    imgPath,
    }

    err = notification.Push()
    if err != nil {
        fmt.Printf("Error showcasing notification: %v", err)
    }
}
```
- **Dynamic Path Handling**: Uses `os.Getwd` and `filepath.Join` to resolve the relative path for the notification icon.
- **Notification Structure**:
  - `AppID`: Identifier for the notification.
  - `Title` and `Message`: Descriptive text displayed in the notification.
  - `Icon`: Path to an image file (must exist).
- **Error Handling**: Logs errors to the console if notifications fail.

---

## **File Structure**
Ensure the following structure for the project:
```
project-root/
├── go.mod
├── go.sum
├── main.go
├── assets/
│   └── clock.png
```
- **`assets/clock.png`**: The icon displayed in the notification. Replace with your custom image if needed.

---

## **Setup and Execution**
1. Place the `clock.png` icon in the `assets` directory.
2. Run the program using:
   ```bash
   go run main.go
   ```
3. Ensure that notifications are enabled in Windows settings:
   - Go to **Settings > System > Notifications & actions**.
   - Enable notifications for the app (listed as "Pomodoro Clock").

---

## **Customizations**
- **Durations**: Modify the constants for work, short break, and long break durations.
- **Icon**: Replace `clock.png` with another image by updating the `assets` directory.
- **Notification Text**: Adjust messages passed to `showNotification` for personalized prompts.

---

## **Common Issues and Fixes**
1. **Notifications Not Displaying**:
   - Verify that the `clock.png` file exists and the path is correct.
   - Check Windows notifications settings.
2. **Path Errors**:
   - Ensure the project is executed from the root directory or use the dynamic path resolution provided.
